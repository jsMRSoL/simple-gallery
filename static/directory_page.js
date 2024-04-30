const imgs = document.querySelectorAll("img");
const images = Array.from(imgs).map((im) => {
  return im.dataset.target;
});
const image_count = images.length;

// console.log(images);
// console.log(image_count);
const displayCanvas = document.querySelector(".media-container");
displayCanvas.addEventListener(
  "touchstart",
  function(event) {
    touchstartX = event.changedTouches[0].screenX;
  },
  false,
);
displayCanvas.addEventListener(
  "touchend",
  function(event) {
    touchendX = event.changedTouches[0].screenX;
    handleGesture();
  },
  false,
);

page = document.querySelector("html");
page.addEventListener("keyup", function(event) {
  if (!displayCanvas.classList.contains("visible")) {
    return;
  }
  switch (event.key) {
    case "n":
      next();
      break;
    case "p":
      prev();
      break;
    case "q":
      unload();
  }
});

const handleGesture = () => {
  if (Math.abs(touchendX - touchstartX) < 100) {
    return;
  }
  if (touchendX < touchstartX) {
    next();
  }
  if (touchstartX < touchendX) {
    prev();
  }
};

const displayItem = (event) => {
  const target = event.target;
  const dt = target.dataset.target;
  load(dt);
  return;
};

const isImage = (filePath) => {
  const imageTypes = ["jpg", "gif", "jpeg", "webp", "png"];
  const videoTypes = ["mp4", "mkv", "webm", "mov", "ogg"];
  const ext = filePath.split(".").pop().toLowerCase();

  if (imageTypes.includes(ext)) {
    return true;
  }
  return false;
};

const load = (filePath) => {
  const mediaContainer = document.querySelector(".media-container");
  mediaContainer.classList.add("visible");
  const mediaItem = document.querySelector(".display-item");
  if (mediaItem != null) {
    mediaContainer.removeChild(mediaItem);
  }
  displayNode = null;
  if (isImage(filePath)) {
    displayNode = document.createElement("img");
  } else {
    displayNode = document.createElement("video");
    displayNode.controls = true;
    displayNode.autoplay = true;
    displayNode.muted = true;
    displayNode.loop = true;
  }
  displayNode.src = filePath;
  displayNode.classList.add("display-item");

  mediaContainer.appendChild(displayNode);
  return;
};

const unload = () => {
  const mediaContainer = document.querySelector(".media-container");
  const mediaItem = document.querySelector(".display-item");
  if (mediaItem != null) {
    mediaContainer.removeChild(mediaItem);
  }
  mediaContainer.classList.remove("visible");
  return;
};

const next = () => {
  const mediaItem = document.querySelector(".display-item");
  const curr_src = mediaItem.src;
  const url = curr_src.replace(window.location.origin, "");
  const filePath = url.replace(/%20/g, " ");
  const idx_current = images.indexOf(filePath);
  next_image = images[idx_current + 1];
  if (next_image === undefined) {
    load(images[0]);
    return;
  }
  load(next_image);
  return;
};

const prev = () => {
  const mediaItem = document.querySelector(".display-item");
  const curr_src = mediaItem.src;
  const url = curr_src.replace(window.location.origin, "");
  const filePath = url.replace(/%20/g, " ");
  const idx_current = images.indexOf(filePath);
  var idx_prev = idx_current - 1;
  if (idx_prev < 0) {
    var idx_last = image_count;
    while (images[idx_last] === undefined) {
      idx_last = idx_last - 1;
    }
    load(images[idx_last]);
    return;
  }
  load(images[idx_prev]);
  return;
};
