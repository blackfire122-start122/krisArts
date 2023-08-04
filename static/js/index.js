const ulElement = document.querySelector('.listArts');
const findInput = document.querySelector('.findInput');
const spanFind = document.querySelector('.spanFind');

const line1 = document.querySelector('.line-1');
const line2 = document.querySelector('.line-2');
const line3 = document.querySelector('.line-3');

line1.scroll()

line2.scrollBy(400, 0);
line3.scrollBy(100, 0);

function scrollHorizontally() {
  line1.scrollBy(1, 0);
  line2.scrollBy(1, 0);
  line3.scrollBy(1, 0);

  const maxScrollLeft = line1.scrollWidth - line1.clientWidth-1;


  if (line1.scrollLeft >= maxScrollLeft) {
    line1.scrollLeft -= line1.firstElementChild.width + 1
    line1.appendChild(line1.firstElementChild)
  }

  if (line2.scrollLeft >= maxScrollLeft) {
    line2.scrollLeft -= line2.firstElementChild.width + 1
    line2.appendChild(line2.firstElementChild)
  }

  if (line3.scrollLeft >= maxScrollLeft) {
    line3.scrollLeft -= line3.firstElementChild.width + 1
    line3.appendChild(line3.firstElementChild)
  }
}

setInterval(scrollHorizontally, 10);

let countArts = 60

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
    <a href="/art/${artData.ID}">
    <h3>${artData.Name}</h3>
    <img src="${artData.Image}" alt="Artwork" width="200" height="200">
    <p>Description: ${artData.Description}</p>
    <p>Price: ${artData.Price}</p>
    <\a>
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
