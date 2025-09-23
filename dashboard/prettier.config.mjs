const config = {
  trailingComma: "es5",
  printWidth: 80,
  tabWidth: 2,
  useTabs: false,
  semi: true,
  singleQuote: false,
  bracketSpacing: true,
  bracketSameLine: false,
  arrowParens: "always",
  plugins: ["prettier-plugin-tailwindcss"],
  overrides: [
    {
      files: "*.svelte",
      options: {
        parser: "svelte",
      },
    },
  ],
  tailwindStylesheet: "./src/app.css",
};

export default config;
