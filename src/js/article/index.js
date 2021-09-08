document.addEventListener('DOMContentLoaded', () => {
  const deleteBtns = document.querySelectorAll('.articles__item-delete');

  const deleteArticle = id => {
    let statusCode;

    fetch(`/${id}`, {
      method: 'DELETE',
    })
      .then(res => {
        statusCode = res.status;
        return res.json();
      })
      .then(data => {
        console.log(JSON.stringify(data));
        if (statusCode == 200) {
          document.querySelector(`.articles__item-${id}`).remove();
        }
      })
      .catch(err => console.error(err));
  };

  // 削除ボタンそれぞれに対してイベントリスナーを設定
  for (let elm of deleteBtns) {
    elm.addEventListener('click', event => {
      event.preventDefault();

      deleteArticle(elm.dataset.id);
    });
  }
});
