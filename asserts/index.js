const app = {
    data() {
        return {
            show_main: false,
            show_input: true,
            add_password: false,
            }
        },
    methods: {
        dark_checked() {
            const check_box = document.querySelector('.dark-theme');

            check_box.addEventListener('click', function(e) {
                if (this.checked) {
                    document.body.classList.toggle('dark-theme');
                } else {
                    document.body.classList.toggle('light-theme');
                }
            })
    }
}
}

Vue.createApp(app).mount('#app-main')
