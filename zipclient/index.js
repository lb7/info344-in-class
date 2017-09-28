$(document).ready(function() {
   $('#button-submit').click(submitName);
   setInterval(getMemory, 1000);
});

function submitName() {
    var name = $('#input-name').val();

    $.get('http://localhost:4000/hello', {name: name}, function (data) {
        $('#response').text(data);
    });
}

function getMemory() {
    $.get(
        'http://localhost:4000/memory',
        {},
        function (data) {
            $('#memory').text(data.Alloc);
        },
        'json'
    );
}