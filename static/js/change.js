const fileInput = document.getElementById('image');
const fileBtn = document.getElementById('file-btn');
const image = document.getElementById('artImage');

fileInput.addEventListener('change',function() {
  const file = fileInput.files[0];
  if (file) {
    const imageUrl = URL.createObjectURL(file);
    image.src = imageUrl;
    fileBtn.innerText = file.name;
  } else {
    image.src = '';
    fileBtn.innerText = 'Файл не вибрано';
  }
});

fileBtn.addEventListener('click', function() {
    fileInput.click();
  });