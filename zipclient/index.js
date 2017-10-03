$(document).ready(function() {
   $('#form-zip').submit(submitForm)
});

function submitForm(e) {
    e.preventDefault();
    
    var city = $('#input-city').val();
    
    $.get('http://localhost:4000/zips/' + city, {}, function (data) {
        var $selectStates = $('#select-states');
        var $selectZips = $('#select-zips');

        var states = [];

        $.each(data, function(idx, val) {
            if (states.indexOf(val.State > -1) {
                states.push(val.State);
            }
        });
    });
}
