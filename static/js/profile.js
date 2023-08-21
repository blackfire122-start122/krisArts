const artImages = document.querySelectorAll(".artImg");
const arts = document.querySelector(".arts");
const fullscreenImageContainer = document.getElementById("fullscreenImageContainer");
const fullscreenImage = document.getElementById("fullscreenImage");
const closeButton = document.getElementById("closeButton");

let imgSelect

function artClick (e) {
    fullscreenImage.src = e.src;
    fullscreenImageContainer.style.display = "block";
    imgSelect = e
}

artImages.forEach((img) => {
    img.addEventListener("click", ()=>artClick(img))
})

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

function goToUrlChange(){
    window.location.href = "/change/"+imgSelect.id
}

let countArts = 20

function findArts() {
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                let response = JSON.parse(xhr.responseText);

                if (response === null){
                    return
                }

                response.forEach(artData => {
                    const art = createArtElement(artData);
                    arts.appendChild(art);
                });
                countArts += 20
            } else {
                console.error('Помилка ' + xhr.status + ': ' + xhr.statusText);
            }
        }
    };
    const url = `http://localhost:8080/api/profile/loadArtsUser?countArts=${encodeURIComponent(countArts)}`;

    xhr.open('GET', url, true);
    xhr.send();
}

function createArtElement(artData) {
    const art = document.createElement('img');

    art.src = artData.Image
    art.alt = "Artwork"
    art.id = artData.ID
    art.className = "artImg"

    art.addEventListener("click", ()=>artClick(art))

    return art;
}

let timerId
let canFindArts = true

window.addEventListener('scroll', handleScroll);

function handleScroll() {
    const contentHeight = document.documentElement.scrollHeight;
    const visibleHeight = window.innerHeight;
    const scrolledHeight = window.scrollY;

    if (contentHeight - (visibleHeight + scrolledHeight) <= 300 && canFindArts ){
        canFindArts = false
        timerId = setTimeout(()=>{canFindArts=true}, 500);
        findArts()
    }
}

function BasketClick(){
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/getAllArtsBasket", true);

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                let data = JSON.parse(xhr.response)

                for (let art in data) {
                    console.log(data[art])
                }

            } else {
                console.log(xhr.statusText);
            }
        }
    };

    xhr.send();
}