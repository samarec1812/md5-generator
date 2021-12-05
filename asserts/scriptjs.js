function Checked(f) {
    $("#sending-file").submit(function (){
        var password = document.getElementById("inputPassword")
        var confirm_pass = document.getElementById("inputPasswordConfirm")
        if (!password && !confirm_pass) {
            confirm("Click OK to sending?");
            f.submit();
        }
        if (password.value !== confirm_pass.value) {
            alert('The passwords are different!')
        } else {
            confirm("Click OK to sending?");
            f.submit();
        }
    })
}

function CheckedText(f) {
    $("#sending-text").submit(function (){
        var password = document.getElementById("inputPasswordText")
        var confirm_pass = document.getElementById("inputPasswordConfirmText")
        if (!password && !confirm_pass) {
            confirm("Click OK to sending?");
            f.submit();
        }
        if (password.value !== confirm_pass.value) {
            alert('The passwords are different!')
        } else {
            confirm("Click OK to sending?");
            f.submit();
        }
    })
}

    // if (password.length != 0 && confirm.length != 0) {
    //         $('#submit-button-file').removeAttr('disabled')
    // } else {
    //     $('#submit-button-file').attr('disabled', 'disabled')
    // }
// }