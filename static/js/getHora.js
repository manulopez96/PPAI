function updateClock() {
    fetch("/time")
      .then((res) => res.json())
      .then((data) => {
        document.getElementById("clock").textContent =
          "Hora actual: " + data.time;
      });
  }

  setInterval(updateClock, 1000);
  updateClock();