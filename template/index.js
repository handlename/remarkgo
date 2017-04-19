document.addEventListener('DOMContentLoaded', function() {
    let slideshow,
        markdownPath;

    function load() {
        slideshow = remark.create({ sourceUrl: markdownPath });
    }

    function reset() {
        slideshow.pause();

        document.getElementsByTagName("body").forEach(function(elem) {
            elem.classList.remove("remark-container");
        });

        document.getElementsByTagName("html").forEach(function(elem) {
            elem.classList.remove("remark-container");
        });

        [
            "remark-slides-area",
            "remark-notes-area",
            "remark-preview-area",
            "remark-backdrop",
            "remark-pause",
            "remark-help"
        ].forEach(function(className) {
            document.getElementsByClassName(className).forEach(function(elem) {
                elem.remove();
            });
        });
    }

    function initSrcPath() {
        markdownPath = document.getElementById("markdownPath").value;
    }

    function initEvents() {
        document.addEventListener("keyup", function (e) {
            if (e.key === "r") {
                reset();
                load();
            }
        });
    }

    function init() {
        initSrcPath();
        initEvents();
        load();
    }

    init();
});
