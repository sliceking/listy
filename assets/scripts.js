 function drag() {
    dragula([
        $('#backlog')[0], 
        $('#working')[0], 
        $('#done')[0]]
    );
 }
 
 $(document).ready(function(){
    drag();
 })

 document.addEventListener("up:modal:close", function(event) {
    console.log('modal closed', event)
    // jumps the history back twice when the modal (show page) is 
    // closed so it doesnt get caught with extra backs to fragments
    window.history.go(-1);
})

document.addEventListener("up:modal:closed", function(event) {
    drag();
})