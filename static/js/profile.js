const artImages = document.querySelectorAll(".artImg");
const arts = document.querySelector(".arts");
const fullscreenImageContainer = document.getElementById("fullscreenImageContainer");
const fullscreenImage = document.getElementById("fullscreenImage");
const closeButton = document.getElementById("closeButton");

const closeButtonOrder = document.getElementById("closeButtonOrder");

const basketPopup = document.querySelector(".basketPopup")
const basketArts = document.querySelector(".basketArts")

const order = document.querySelector(".order")

const closeButtonBasket = document.getElementById("closeButtonBasket");

let imgSelect

function artClick (e) {
    closeBasket()
    fullscreenImage.src = e.src;
    fullscreenImageContainer.style.display = "block";
    imgSelect = e
}

artImages.forEach((img) => {
    img.addEventListener("click", ()=>artClick(img))
})

function closeImage(){
    fullscreenImageContainer.style.display = "none"
}
closeButton.addEventListener("click", closeImage);

closeButtonBasket.addEventListener("click", closeBasket);

closeButtonOrder.addEventListener("click", closeOrder);

function closeBasket(){
    basketPopup.style.display = "none"
}

function closeOrder(){
    order.style.display = "none"
}

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
    closeImage()
    basketPopup.style.display = "flex"
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/getAllArtsBasket", true);

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                let data = JSON.parse(xhr.response)
                basketArts.innerHTML = ""
                for (let i in data) {
                    basketArts.appendChild(artBasket(data[i]))
                }

            } else {
                console.log(xhr.statusText);
            }
        }
    };

    xhr.send();
}

function artBasket(artData){
    let art = document.createElement("div")
    let link = document.createElement("a")
    let name = document.createElement("h2")
    let image = document.createElement("img")
    let price = document.createElement("p")
    let divTexts = document.createElement("div")
    let divBtnPrice = document.createElement("div")
    let deleteBtn = document.createElement("button")

    
    art.className = "artBasket"
    name.innerHTML = artData.Name
    image.src = "/"+artData.Image
    price.innerHTML = artData.Price+"₴"
    deleteBtn.className = "deleteBtn"
    deleteBtn.innerHTML = "Delete"
    divBtnPrice.className = "divBtnPrice"
    divTexts.className = "divTexts"
    link.href = "/art/"+artData.ID
    link.className = "link"

    deleteBtn.addEventListener("click",()=>{
        let xhr = new XMLHttpRequest();
        let artId = encodeURIComponent(artData.ID);

        xhr.open("DELETE", "/api/deleteFromBasket/" + artId, true);

        xhr.onreadystatechange = function() {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    art.remove()
                } else {
                    console.log(xhr.statusText);
                }
            }
        };

        xhr.send();
    })


    art.append(image)
    art.append(divTexts)

    divTexts.append(link)
    divTexts.append(divBtnPrice)

    link.append(name)

    divBtnPrice.append(deleteBtn)
    divBtnPrice.append(price)

    return art
}

function orderClick(){
  closeBasket()
  order.style.display = "flex"
  
  
}