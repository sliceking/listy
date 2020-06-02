 function setupDragging() {
    var dragContainers = dragula([
        $('#backlog')[0], 
        $('#working')[0], 
        $('#done')[0]]
    );

    dragContainers.on('drop', function(el, target, source){
        console.log('el', el);

        if (target !== source) {
            console.log('different');
        } else {
            console.log('same');
        }
    })
 }
 
 $(document).ready(function(){
    setupDragging();
 })

 document.addEventListener("up:modal:close", function(event) {
    console.log('modal closed', event)
    // jumps the history back twice when the modal (show page) is 
    // closed so it doesnt get caught with extra backs to fragments
    window.history.go(-1);
})

document.addEventListener("up:modal:closed", function(event) {
    setupDragging();
})