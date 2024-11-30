document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('loginForm');
    const usernameInput = document.getElementById('username');
    const passwordInput = document.getElementById('password');

    form.addEventListener('submit', (event) => {
        const username = usernameInput.value.trim();
        const password = passwordInput.value.trim();

        const usernameRegex = /^[a-zA-Z0-9._-]{3,16}$/; // Alphanumeric, 3-16 chars
        const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$/; // Min 8 chars, one upper, one lower, one digit


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
        alert("Validation successful. Logging in...");
    });
});
