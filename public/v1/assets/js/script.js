window.onscroll = () => {
    if ($(window).width() > 768) {
        if (window.scrollY > 40) {
            $(".user-information-wrapper").attr(
            "style",
            "display: none !important; transition: 0.2s;"
            );
            $('#header').addClass('hidden-header');
        } else {
            $(".user-information-wrapper").attr(
            "style",
            "display: block; transition: 0.2s;"
            );
            $('#header').removeClass('hidden-header');
        }
    }
};

function Controlling(condition){
    var conditionSplitter = condition.split('"').join("")
    Swal.fire({
        title: '<span style="font-size: 16px;font-weight: 400;">Apakah Kamu Yakin Ingin Mengubah Kondisi Perangkat?</span>',
        showDenyButton: true,
        confirmButtonText: 'Ya',
        confirmButtonColor: '#282689',
        denyButtonText: `Tidak`,
        }).then((result) => {
        if (result.isConfirmed) {
            window.location = "/control/" + conditionSplitter;
        } else if (result.isDenied) {
            window.location = "/daftar-perangkat";
        }
    })
}