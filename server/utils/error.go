package utils

var ErrorPageContents = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/css/app.css">
    <link rel="stylesheet" href="/assets/css/error.css">
    <link rel="shortcut icon" href="/assets/images/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
</head>

<body>
    <header>
        <a href="/">
            <p>
                <img class="img-01" src="/assets/images/01.png" alt="01">
                <span class="forum-title">forum</span>
            </p>
        </a>
    </header>
    <div class="error-page">
        <div class="error-card">
            <h1>500</h1>
            <p>Internal Server Error</p>
            <a href="/"><i class="fa-solid fa-circle-left"></i> Back Home</a>
        </div>
    </div>
    <div class="footer-container">
        <footer>
            <div>
                <span>01</span>Forum copyright © 2024
            </div>
            <div class="footer-team">
                <a href="https://github.com/hmaach/"><img src="/assets/images/github.png" alt="github-logo" /> Hamza
                    Maach</a> |
                <a href="https://github.com/ABouziani"><img src="/assets/images/github.png" alt="github-logo" />
                    Abdelhamid
                    Bouziani</a> |
                <a href="https://github.com/2001basta"><img src="/assets/images/github.png" alt="github-logo" /> Youssef
                    Basta</a> |
                <a href="https://github.com/oaitbenh"><img src="/assets/images/github.png" alt="github-logo" /> Omar Ait
                    Benhammou</a> |
                <a href="https://github.com/M-MDI"><img src="/assets/images/github.png" alt="github-logo" /> Mehdi
                    Moulabbi</a>
            </div>
        </footer>
    </div>
    <script src="/assets/js/index.js"></script>

</body>

</html>
`