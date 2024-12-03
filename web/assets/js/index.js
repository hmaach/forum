function throttle(fn, delay) {
    let last = 0;
    return function () {
        const now = +new Date();
        if (now - last > delay) {
            fn.apply(this, arguments);
            last = now;
        }
    };
}

const addcomment = throttle(addcomm, 5000)

function postreaction(postId, reaction) {
    document.getElementById("errorlogin" + postId).innerText = ``
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/postreaction", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            document.getElementById("likescount" + postId).innerHTML = `<i
                    class="fa-regular fa-thumbs-up"></i>${response.likesCount}`;
            document.getElementById("dislikescount" + postId).innerHTML = `<i
                    class="fa-regular fa-thumbs-down"></i>${response.dislikesCount}`;
        } else if (xhr.status !== 200) {
            document.getElementById("errorlogin" + postId).innerText = `You must login first!`
            setTimeout(() => {
                document.getElementById("errorlogin" + postId).innerText = ``
            }, 1000);
        }
    };
    xhr.send(`reaction=${reaction}&post_id=${postId}`);
}
function commentreaction(commentid, reaction) {
    document.getElementById("commenterrorlogin" + commentid).innerText = ``
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/commentreaction", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            document.getElementById("commentlikescount" + commentid).innerHTML = `<i
                    class="fa-regular fa-thumbs-up"></i>${response.commentlikesCount}`;
            document.getElementById("commentdislikescount" + commentid).innerHTML = `<i
                    class="fa-regular fa-thumbs-down"></i>${response.commentdislikesCount}`;
        } else if (xhr.status !== 200) {
            document.getElementById("commenterrorlogin" + commentid).innerText = `You must login first!`
            setTimeout(() => {
                document.getElementById("commenterrorlogin" + commentid).innerText = ``
            }, 1000);

        }
    };
    xhr.send(`reaction=${reaction}&comment_id=${commentid}`);
}


function addcomment(postId) {
    const content = document.getElementById("comment-content");
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/addcommentREQ", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            const comment = document.createElement("div")
            comment.innerHTML = `
                 <div class="comment">
            <div class="comment-header">
                <p class="comment-user">`+ response.username + `</p>
                <span></span>
                <p class="comment-time">`+ response.created_at + ` </p>
            </div>
            <div class="comment-body">
                <p class="comment-content">`+ response.content + ` </p>
            </div>
            <div class="comment-footer">
                <button id="commentlikescount`+ response.ID + `" onclick="commentreaction('` + response.ID + `','like')"
                    class="comment-like"><i class="fa-regular fa-thumbs-up"></i>`+ response.likes + `</button>
                <button id="commentdislikescount`+ response.ID + `" onclick="commentreaction('` + response.ID + `','dislike')"
                    class="comment-dislike"><i class="fa-regular fa-thumbs-down"></i>`+ response.dislikes + `</button>
            </div>
            <span style="color:red" id="commenterrorlogin`+ response.ID + `"></span>
        </div>
                `
            document.getElementsByClassName("comments")[0].prepend(comment)
            document.getElementsByClassName("post-comments")[0].innerHTML = `<i class="fa-regular fa-comment"></i>` + response.commentscount
            content.value = ""
        } else if (xhr.status === 400) {
            document.getElementById("errorlogin" + postId).innerText = `Invalid comment!`
            setTimeout(() => {
                document.getElementById("errorlogin" + postId).innerText = ``
            }, 1000);
        } else if (xhr.status === 401) {
            document.getElementById("errorlogin" + postId).innerText = `You must login first!`
            setTimeout(() => {
                document.getElementById("errorlogin" + postId).innerText = ``
            }, 1000);
        }
    };
    xhr.send(`postid=${postId}&comment=${content.value}`);
}

const select = document.getElementById('categories-select');
if (select) {

    select.addEventListener('change', (e) => {
        // Parse the value as JSON to extract id and label
        const selectedValue = JSON.parse(e.target.value);
        const { id, label } = selectedValue;
        // console.log('ID:', id, 'Label:', label);

        // create the elemenet for the category
        const span = document.createElement('span');
        span.textContent = label;
        span.classList.add('selected-category');


        // create hidden input to hold the id of selected category
        const input = document.createElement('input')
        input.type = 'hidden';
        input.value = id
        input.name = 'categories'

        // add the elements (span and hidden input) 
        // at the first  position of the categories container
        const categoriesContainer = document.querySelector('.selected-categories');
        categoriesContainer.append(input, span);


        // disable the option selected in the select
        e.target.options[e.target.selectedIndex].disabled = true;

        // Reset the select 
        e.target.selectedIndex = 0;
    });

}

async function pagination(dir, data) {
    const path = window.location.pathname
    if (dir === "next" && data) {
        const page = +document.querySelector(".currentpage").innerText + 1
        window.location.href = path + "?PageID=" + page;
    }

    if (dir === "back" && document.querySelector(".currentpage").innerText > "1") {
        const page = +document.querySelector(".currentpage").innerText - 1
        window.location.href = path + "?PageID=" + page;
    }
}
