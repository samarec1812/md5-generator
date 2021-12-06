var checkedEmpty = false
function Checked(f) {
    $("#sending-file").submit(function (){
        var password = document.getElementById("inputPassword")
        var confirm_pass = document.getElementById("inputPasswordConfirm")
        var complexity = document.getElementById("complexity1")
        var len_password = document.getElementById("key-len1")
        if (!password && !confirm_pass && !complexity && !len_password) {
            confirm("Click OK to sending?");
            f.submit();
        } else {

        console.log(complexity.value, len_password.value)
        if (password.value.length != len_password.value) {
               alert('The actual password length does not match the specified!')
               return
        }
        console.log(typeof complexity.value)
            if (complexity.value === "0" && !(/\d/.test(password.value))) {
                alert('Password must consist numbers')
                return
            }
            if (complexity.value === "1" && !(/\d/.test(password.value) && /[a-zA-Z]/.test(password.value))) {
                alert('Password must consist of upper and lower case letters of the Latin alphabet and numbers')
                return
            }
            if (complexity.value === "2" && !(/\d/.test(password.value) && /[a-zA-Z]/.test(password.value) && /[а-яА-Я]/.test(password.value))) {
                alert('Password must consist of upper and lower case letters of the Latin alphabet and Cyrillic alphabet, as well as numbers')
                return
            }
            if (complexity.value === "3" && !(/\d/.test(password.value) && /[a-zA-Z]/.test(password.value) && /[а-яА-Я]/.test(password.value) && /[ !@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password.value))) {
                alert('Password must consist of upper and lower case letters of the Latin alphabet and Cyrillic alphabet, as well as numbers and special characters')
                return
            }

        if (password.value !== confirm_pass.value) {
            alert('The passwords are different!')


        } else {
            confirm("Click OK to sending?");
            f.submit();
        }

        }
    })
}

function CheckedText(f) {
    $("#sending-text").submit(function (){
        var password = document.getElementById("inputPasswordText")
        var confirm_pass = document.getElementById("inputPasswordConfirmText")
        var complexity = document.getElementById("complexity2")
        var len_password = document.getElementById("key-len2")
        if (checkedEmpty) {
            var emptyText = document.getElementById("floatingTextarea")
            if (emptyText.value !== "@") {
                alert('The text data is incorrect. Enter \'@\' to hash an empty string')
                return
            }
        }

        if (!password && !confirm_pass && !complexity && !len_password) {
            confirm("Click OK to sending?");
            f.submit();
        } else {

        console.log(complexity.value, len_password.value)
        if (password.value.length != len_password.value) {
            alert('The actual password length does not match the specified!')
            return
        }
        console.log(typeof complexity.value)
        if (complexity.value === "0" && !(/\d/.test(password.value))) {
            alert('Password must consist numbers')
            return
        }
        if (complexity.value === "1" && !(/\d/.test(password.value) && /[a-zA-Z]/.test(password.value))) {
            alert('Password must consist of upper and lower case letters of the Latin alphabet and numbers')
            return
        }
        if (complexity.value === "2" && !(/\d/.test(password.value) && /[a-zA-Z]/.test(password.value) && /[а-яА-Я]/.test(password.value))) {
            alert('Password must consist of upper and lower case letters of the Latin alphabet and Cyrillic alphabet, as well as numbers')
            return
        }
        if (complexity.value === "3" && !(/\d/.test(password.value) && /[a-zA-Z]/.test(password.value) && /[а-яА-Я]/.test(password.value) && /[ !@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password.value))) {
            alert('Password must consist of upper and lower case letters of the Latin alphabet and Cyrillic alphabet, as well as numbers and special characters')
            return
        }

        if (password.value !== confirm_pass.value) {
            alert('The passwords are different!')
        } else {
            confirm("Click OK to sending?");
            f.submit();

        }}
    })
}


function ChangeEmptyText() {
    checkedEmpty = !checkedEmpty
    return checkedEmpty
}
    // if (password.length != 0 && confirm.length != 0) {
    //         $('#submit-button-file').removeAttr('disabled')
    // } else {
    //     $('#submit-button-file').attr('disabled', 'disabled')
    // }
// }