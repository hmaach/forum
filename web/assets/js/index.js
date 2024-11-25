function postreaction(postId, reaction) {
    document.getElementById("errorlogin"+postId).innerText = ``
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/post/postreaction", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            document.getElementById("likescount"+postId).innerHTML = `<i
                    class="fa-regular fa-thumbs-up"></i>${response.likesCount}`;
            document.getElementById("dislikescount"+postId).innerHTML = `<i
                    class="fa-regular fa-thumbs-down"></i>${response.dislikesCount}`;
        } else if (xhr.status !== 200) {
            document.getElementById("errorlogin"+postId).innerText = `You must login first!`
            setTimeout(() => {
                document.getElementById("errorlogin"+postId).innerText = ``
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
