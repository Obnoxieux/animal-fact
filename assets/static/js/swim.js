// Select the duck and pond elements
const duck = document.getElementById('duck');
const pond = document.querySelector('.pond');

if (duck && pond) {
    // Get the pond's boundaries
    const pondRect = pond.getBoundingClientRect();

    // Set initial duck position and angle
    let position = {x: pondRect.width / 2, y: pondRect.height / 2}; // Start in the center
    let angle = Math.random() * 360; // Random initial angle

    // Set speed of the duck
    const speed = 1.2; // Pixels per frame

    // Function to move the duck
    function swim() {
        // Calculate new position based on angle and speed
        position.x += Math.cos((angle * Math.PI) / 180) * speed;
        position.y += Math.sin((angle * Math.PI) / 180) * speed;

        // Check if the duck hits the edges of the pond
        if (position.x < 0) {
            position.x = 0;
            angle = 180 - angle; // Reverse horizontal direction
        } else if (position.x > pondRect.width - duck.offsetWidth) {
            position.x = pondRect.width - duck.offsetWidth;
            angle = 180 - angle; // Reverse horizontal direction
        }
        if (position.y < 0) {
            position.y = 0;
            angle = 360 - angle; // Reverse vertical direction
        } else if (position.y > pondRect.height - duck.offsetHeight) {
            position.y = pondRect.height - duck.offsetHeight;
            angle = 360 - angle; // Reverse vertical direction
        }

        // Ensure angle stays within 0-360 degrees
        angle = angle % 360;

        // Update duck position
        duck.style.transform = `translate(${position.x}px, ${position.y}px) rotate(${angle}deg)`;

        // Repeat animation
        requestAnimationFrame(swim);
    }

    // Set initial duck styles
    duck.style.position = 'absolute';
    duck.style.left = `${pondRect.left}px`;
    duck.style.top = `${pondRect.top}px`;

    // Start the animation
    swim();
}
