/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      spacing:{
        text: "clamp(45ch,50%,75ch)"
      }
    },
  },
  plugins: [],
}

