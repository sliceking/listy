 function setupDragging() {
    var dragContainers = dragula([
        $('#backlog')[0], 
        $('#working')[0], 
        $('#done')[0]]
    );

    dragContainers.on('drop', function(el, target, source){
        console.log('el', el.dataset.id);
        console.log(target.id);

        if (target !== source) {
            const data = {category: target.id};
            const options = {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data),              
            };

            const request = new Request(`/cards/${el.dataset.id}`, options)
            fetch(request)
                .then(response => console.log(response));

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