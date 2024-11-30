document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('registerForm');
    const emailInput = document.getElementById('email');
    const usernameInput = document.getElementById('username');
    const passwordInput = document.getElementById('password');
    const passwordConfirmationInput = document.getElementById('passwordConfirmation');

    form.addEventListener('submit', (event) => {
        const email = emailInput.value.trim();
        const username = usernameInput.value.trim();
        const password = passwordInput.value.trim();
        const passwordConfirmation = passwordConfirmationInput.value.trim();

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; 
        const usernameRegex = /^[a-zA-Z0-9._-]{3,16}$/;
        const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$/; // Min 8 chars, one upper, one lower, one digit

        if (!emailRegex.test(email)) {
            alert("Invalid email address.");
            event.preventDefault(); 
            return;
        }

        if (!usernameRegex.test(username)) {
            alert("Invalid username. Must be 3-16 characters long and include only letters, numbers, dots, underscores, or dashes.");
            event.preventDefault(); 
            return;
        }

        if (!passwordRegex.test(password)) {
            alert("Invalid password. Must be at least 8 characters long, include one uppercase letter, one lowercase letter, and one number.");
            event.preventDefault(); 
            return;
        }
        
        if (password !== passwordConfirmation) {
            alert("Passwords do not match.");
            event.preventDefault();
            return;
        }

        alert("Registration successful!");
    });
});
