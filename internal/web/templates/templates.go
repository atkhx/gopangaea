package templates

import (
	"bytes"
	"html/template"
	"io"
	"strings"

	"github.com/pkg/errors"
)

func New(root string, paths ...string) (*templates, error) {
	res := &templates{
		root:  root,
		paths: paths,
	}

	if err := res.parse(); err != nil {
		return nil, err
	}

	return res, nil
}

type templates struct {
	root  string
	paths []string
	tmpls *template.Template
}

func (t *templates) parse() (err error) {
	//if t.tmpls == nil {
	t.tmpls = template.New("templates")
	for i := 0; i < len(t.paths); i++ {
		_, err = t.tmpls.ParseGlob(t.root + t.paths[i])
		if err != nil {
			break
		}
	}
	//}

	return
}

func (t *templates) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	return t.tmpls.ExecuteTemplate(w, name, data)
}

func (t *templates) RenderLayout(w io.Writer, name string, data interface{}) error {
	return t.tmpls.ExecuteTemplate(w, "layout/"+strings.Trim(name, "/"), data)
}

func (t *templates) RenderView(w io.Writer, name string, data interface{}) error {
	return t.tmpls.ExecuteTemplate(w, "views/"+strings.Trim(name, "/"), data)
}

func (t *templates) RenderLayoutWithView(w io.Writer, layout, view string, layoutData, viewData map[string]interface{}) error {
	// for debug
	if err := t.parse(); err != nil {
		return err
	}
	// <- for debug

	buf := bytes.NewBuffer(nil)
	if err := t.RenderView(buf, view, viewData); err != nil {
		return errors.Wrap(err, "can't render view")
	}

	if layoutData == nil {
		layoutData = map[string]interface{}{}
	}

	layoutData["content"] = template.HTML(buf.String())

	buf.Reset()

	if err := t.RenderLayout(buf, layout, layoutData); err != nil {
		return errors.Wrap(err, "can't render layout")
	}

	if _, err := buf.WriteTo(w); err != nil {
		return errors.Wrap(err, "can't flush buffer")
	}

	return nil
}
