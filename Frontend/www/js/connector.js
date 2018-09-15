$(document).ready(function() {
    var myApp = new Framework7();
    var $$ = Dom7;

    myApp.on('test3', function() {
        console.log("test");
        alert("test");
    });
});

function LoadRecipe() {
    /* const request = require('request');

     request('http://127.0.0.1:7777/Ings?h=Beef', { json: true }, (err, res, body) => {
         if (err) { return console.log(err); }
         console.log(body.url);
         console.log(body.explanation);
     });*/
    console.log("test");
    alert("test");
}