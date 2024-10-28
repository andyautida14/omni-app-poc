const htmlElement = document.documentElement;

if (
  localStorage.getItem("mode") === "dark" ||
  (!("mode" in localStorage) &&
    window.matchMedia("(prefers-color-scheme: dark)").matches)
) {
  htmlElement.classList.add("dark");
} else {
  htmlElement.classList.remove("dark");
}

htmlElement.classList.add(localStorage.getItem("theme") || "uk-theme-zinc");
