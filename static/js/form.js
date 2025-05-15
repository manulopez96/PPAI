document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("miFormulario");

  form.addEventListener("submit", (e) => {
    const nombre = document.getElementById("nombre").value;
    console.log("Nombre:", nombre);
    if (nombre.length < 3) {
      e.preventDefault();
      alert("El nombre debe tener al menos 3 caracteres.");
    }
  });
});
