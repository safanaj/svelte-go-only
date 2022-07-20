import "./app.scss";
import App from "./App.svelte";

Object.defineProperty(String.prototype, 'capitalize', {
  value: function() {
    return this.charAt(0).toUpperCase() + this.slice(1).toLowerCase();
  },
  enumerable: false
});

const app = new App({
  target: document.getElementById("app"),
});

export default app;
