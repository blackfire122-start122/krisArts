const artImages = document.querySelectorAll(".artImg");
const fullscreenImageContainer = document.getElementById("fullscreenImageContainer");
const fullscreenImage = document.getElementById("fullscreenImage");
const closeButton = document.getElementById("closeButton");

let imgSelect

artImages.forEach((img) => {
    img.addEventListener("click", () => {
        fullscreenImage.src = img.src;
        fullscreenImageContainer.style.display = "block";
        imgSelect = img
    });
});

function closeImage(){
    fullscreenImageContainer.style.display = "none";
}
closeButton.addEventListener("click", closeImage);

function deleteArt(){
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status !== 200) {
                console.error('Помилка ' + xhr.status + ': ' + xhr.statusText);
                return
            }
            imgSelect.remove()
            closeImage()
            hideDeleteConfirmation()
        }
    };
    const url = `http://localhost:8080/api/deleteArt?id=${encodeURIComponent(imgSelect.id)}`;

    xhr.open('DELETE', url, true);
    xhr.send();
}

function showDeleteConfirmation() {
    const deleteConfirmation = document.getElementById("deleteConfirmation");
    deleteConfirmation.style.display = "block";
}

function hideDeleteConfirmation() {
    const deleteConfirmation = document.getElementById("deleteConfirmation");
    deleteConfirmation.style.display = "none";
}