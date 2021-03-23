function $(el){
    return document.querySelector(el)
}

// Узнаём ID фильма
const currentLink = document.location.href.split('/')
const FILM_ID = currentLink[currentLink.length - 1]
// текущая страница для загрузки комментариев(комменты в api по страницам)
let commentPage = 1


btn = $('#more-btn')
commentsList = $('.comments')

// Ajax запрос новых комментов
function getComments(){
    // Подсчёт загруженных комментов для правильного ajax запроса
    let commentsCount = commentsList.querySelectorAll('.comment').length
    let URL = `http://${HOST}/website/comments?page=${commentPage}&id=${FILM_ID}`

    fetch(URL, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
        }).then(response => {
            if (response.ok){
                commentPage++           
                return response.json()
            }
            else
                console.error('Ошибка получения данных');
        }).then( comments => {    
            console.log(comments.film_comments.length)
                                
            comments.film_comments.map( comment => generateComment(comment))            
        })

}

// генерация HTML кода(добавление новых комментов) 
function generateComment(comment){
    let commentElement = document.createElement('div')
    commentElement.classList.add('comment')
    commentElement.innerHTML = `
        <div class="comment__img_wrapper">
            <img src="/img/comment_img.png" class="comment__img">
        </div>
        <div class="comment__name_wrapper">
            <p class="comment__name">${comment.Name}</p>        
        </div>
        <time class="comment__time">${comment.CreatedAt}</time>
        <div class="comment__text_wrapper">
            <p class="comment__text">${comment.Text }</p>
        </div>
    `
    commentsList.append(commentElement)
}

btn.addEventListener('click', () => {
    getComments()
})

getComments()