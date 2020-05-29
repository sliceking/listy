 dragula([
    document.querySelector('#backlog'), 
    document.querySelector('#working'), 
    document.querySelector('#done')]
);

document.addEventListener("up:modal:closed", function(event) {
    console.log('modal closed', event)
    // jumps the history back twice when the modal (show page) is 
    // closed so it doesnt get caught with extra backs to fragments
    window.history.go(-2);
})