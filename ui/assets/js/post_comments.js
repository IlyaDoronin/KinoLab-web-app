// function $(el){
//     return document.querySelector(el)
// }

// Узнаём id фильма из ссылки
const link = document.location.href.split('/')
const filmID = link[link.length - 1]
const URL = `http://${HOST}/website/postComment?filmID=${filmID}`

btn = $('#send-btn')
form = $('#form')


// Ajax запрос новых комментов
function sendComment(data){    
    fetch(URL, {
        method: 'POST',
        body: data
    })
}


form.addEventListener('submit', e => {
    e.preventDefault()
    let data = new FormData(form)
    let name = form.name.value
    if (name == ''){
        name = 'Аноним'
        data.set("name", name)        
    }
    data.append("filmID", filmID)
    sendComment(data)

    
    let currentDate = new Date()
    time = `${currentDate.getFullYear()}-${currentDate.getMonth()}-${currentDate.getDay()} ${currentDate.getHours()}:${currentDate.getMinutes()}:${currentDate.getSeconds()}`

    // Генерация нового коммента
    let commentElement = document.createElement('div')
    commentElement.classList.add('comment')
    commentElement.innerHTML = `
        <div class="comment__img_wrapper">
            <img src="/img/comment_img.png" class="comment__img">
        </div>
        <div class="comment__name_wrapper">
            <p class="comment__name">${ name }</p>        
        </div>
        <time class="comment__time">${time}</time>
        <div class="comment__text_wrapper">
            <p class="comment__text">${ form.text.value }</p>
        </div>
    `
    commentsList.prepend(commentElement)
})
