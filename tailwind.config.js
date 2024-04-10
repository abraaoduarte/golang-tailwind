/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "internal/templates/**/*.{html,js,templ,go}",
        //  "./templates/common/**/*.{html,js,templ,go}",
    ],
    theme: {
        extend: {},
    },
    plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
