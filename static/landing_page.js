const refresh_btn = document.querySelector(".refresh-btn");

refresh_btn.addEventListener(
  "click",
  function() {
    console.log("Got to here");
    if (right_menu.classList.contains("rightnav-visible")) {
      right_menu.classList.remove("rightnav-visible");
      return;
    }
    right_menu.classList.add("rightnav-visible");
  },
  false,
);
