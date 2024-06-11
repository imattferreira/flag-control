/** @type {import('tailwindcss').Config} */
const config = {
  content: ['./index.html', './node_modules/flowbite/**/*.js', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {}
  },
  plugins: [import('flowbite/plugin')]
}

export default config
