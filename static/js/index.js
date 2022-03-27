$('document').ready(function () {
    $('#controlForm input').change(function () {
        if (this.type === 'checkbox') {
            $('#debug').html('changed input ' + this.id + ':' + this.checked)
            sendChange(this.id, this.checked ? "true" : "false");
        } else {
            $('#debug').html('changed input' + this.id + ':' + this.value)
            $('#'+this.id+'_val').text(this.value);
            sendChange(this.id, this.value);
        }
    })

    $('#controlForm select').change(function () {
        $('#debug').html('changed select' + this.id + ':' + this.value);
        sendChange(this.id, this.value);
    })

    $('#changeMode button').click(function () {
        $('#debug').html('changeMode click ' + this.id + ':' + this.value);
        sendChange(this.id, this.value);
        $('#modeBtnGroup').text($(this).text())
    })

    $('#changePreset button').click(function () {
        $('#debug').html('changePreset click ' + this.id + ':' + this.value);
        sendChange(this.id, this.value);
        $('#presetBtnGroup').text($(this).text())
        document.location.href = '/';
    })

    $('#changeBank button').click(function () {
        $('#debug').html('changeBank click ' + this.id + ':' + this.value);
        sendChange(this.id, this.value);
        $('#bankBtnGroup').text($(this).text())
        document.location.href = '/';
    })

    $('#controlForm input[type=range]').each(function (input) {
        $('#debug').html('' + this.id + ':' + this.value);
        let label=$('label[for='+this.id+']');
        label.css("display", "block")
        label.html(
            '<span">'+label.html()+'</span> ['+
            '<span class="col-sm-3 text-end" id="'+this.id+'_val">'+this.value+'</span>]'
        )
    })

    $('#savePresetBtn').click(function () {
        $.ajax({
            url: '/save-preset',
            type: 'POST',
            success: function () {
                document.location.href = '/'
            }
        });
    })

    $('#resetPresetBtn').click(function () {
        $.ajax({
            url: '/reset-preset',
            type: 'POST',
            success: function () {
                document.location.href = '/'
            }
        });
    })
});

function sendChange(name, value) {
    let data = {'params': [{'name': name, 'value': value}]};

    $.ajax({
        url: '/change',
        type: 'POST',
        data: JSON.stringify(data),
        contentType: 'application/json; charset=utf-8',
        dataType: 'text',
        async: false,
        success: function (response) {
            $('#debug').html('changed param ' + name + ':' + value + '; response: ' + response);
        }
    });
}