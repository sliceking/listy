 dragula([
    document.querySelector('#backlog'), 
    document.querySelector('#working'), 
    document.querySelector('#done')]
);

$('#backlog, #working, #done').on('click', detailView);

function detailView(event) {
    console.log('clicked', event);
}