var $$ = Dom7;
$$('.mealtime_test .submit_ing_btn2').on('click', function(e) {
    console.log("test");
});


function LoadRecipe() {
    const request = require('request');

    request('http://127.0.0.1:7777/Ings?h=Beef', { json: true }, (err, res, body) => {
        if (err) { return console.log(err); }
        console.log(body.url);
        console.log(body.explanation);
    });
}
$(document).ready(LoadRecipe)