const ulElement = document.querySelector('.listArts');
const findInput = document.querySelector('.findInput');
const spanFind = document.querySelector('.spanFind');

let countArts = 20

function findArts(e) {
  spanFind.innerText = ""

  let xhr = new XMLHttpRequest();

  if (e){
    countArts=0
  }

  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4) {
      if (xhr.status === 200) {
        let response = JSON.parse(xhr.responseText);

        if (e){
          ulElement.innerHTML = '';
        }

        if (response === null){
          spanFind.innerText = "Nothing not find"
          return
        }

        response.forEach(artData => {
          const liElement = createArtElement(artData);
          ulElement.appendChild(liElement);
        });
        countArts += 20
      } else {
        console.error('Помилка ' + xhr.status + ': ' + xhr.statusText);
      }
    }
  };
  const url = `http://localhost:8080/api/findArts?find=${encodeURIComponent(findInput.value)}&countArts=${encodeURIComponent(countArts)}`;

  xhr.open('GET', url, true);
  xhr.send();
}

function createArtElement(artData) {
  const liElement = document.createElement('li');
  
  liElement.innerHTML = `
    <h3>${artData.Name}</h3>
    <img src="${artData.Image}" alt="Artwork" width="200" height="200">
    <p>Description: ${artData.Description}</p>
    <p>Price: ${artData.Price}</p>
  `;

  return liElement;
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
