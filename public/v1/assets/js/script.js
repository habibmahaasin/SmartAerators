window.onscroll = () => {
    if ($(window).width() > 768) {
        if (window.scrollY > 40) {
            $(".user-information-wrapper").attr(
            "style",
            "display: none !important; transition: 0.2s;"
            );
        } else {
            $(".user-information-wrapper").attr(
            "style",
            "display: block; transition: 0.2s;"
            );
        }
    }
};
