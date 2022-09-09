<button id="open-modal">Open Modal</button>

<div class="modal">
  <div class="modal-content">
    <div class="modal-header">
      <span id="close-modal">&times;</span>
      <h3>Simple Modal</h3>
    </div>
    <div class="modal-body">
      <p>Some text</p>
    </div>
  </div>
</div>

<style>
  @import url("https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700");

  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: "Open Sans", sans-serif;
    min-height: 100vh;
  }

  /* Open modal button styling */
  #open-modal {
    border: none;
    background-color: #2196f3;
    border-radius: 1.5rem;
    padding: 0.7rem 2.3rem;
    font-size: 1.1rem;
    font-weight: bold;
    cursor: pointer;
    color: #fff;
    user-select: none;
  }

  #open-modal:hover {
    filter: brightness(1.1);
  }

  /* Modal styling */
  .modal {
    display: none;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.2);
  }

  .modal-content {
    margin: 20% auto;
    max-width: 500px;
    background-color: #fff;
    padding: 20px;
    animation: scale 0.5s ease;
    border-radius: 10px;
    box-shadow: 6.7px 6.7px 5.3px rgba(0, 0, 0, 0.02),
      22.3px 22.3px 17.9px rgba(0, 0, 0, 0.03),
      100px 100px 80px rgba(0, 0, 0, 0.05);
  }

  .modal-body {
    margin-top: 5px;
  }

  #close-modal {
    float: right;
    font-size: 1.2rem;
    font-weight: bold;
    cursor: pointer;
  }

  /* Modal animation */
  @keyframes scale {
    from {
      transform: scale(0);
    }

    to {
      transform: scale(1);
    }
  }

</style>

<style>const modal = document.querySelector(".modal");
const openModalBtn = document.querySelector("#open-modal");
const closeModalBtn = document.querySelector("#close-modal");

openModalBtn.addEventListener("click", function () {
  modal.style.display = "block";
});

closeModalBtn.addEventListener("click", function () {
  modal.style.display = "none";
});
</style>

<script>
  const modal = document.querySelector(".modal");
 const openModalBtn = document.querySelector("#open-modal");
 const closeModalBtn = document.querySelector("#close-modal");

 openModalBtn.addEventListener("click", function () {
   modal.style.display = "block";
 });

 closeModalBtn.addEventListener("click", function () {
   modal.style.display = "none";
 });

</script>
