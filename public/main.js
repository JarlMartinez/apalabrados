const d = document;

const charactersContainer = d.getElementById("characters-list")
const textsContainer = d.getElementById("texts-list")
const numbersContainer = d.getElementById("numbers-list")
const inputField = d.getElementById("user-input")
const submitInputButton = d.getElementById("submit-input")
const errorP = d.getElementById("error-message")

submitInputButton.addEventListener("click", () => {
  const data = inputField.value
  if (!data || data.trim() === "") {
    return errorP = 'Input must have content';
  }
  postInput(data)
  .then(() => refresh())
  .catch(e => errorP = 'Error posting input.' + e)
})

// Initiate UI
refresh()

function refresh() {
  inputField.value = ""
  refreshNumbersList()
  refreshTextsList()
  refreshCharactersList()
}

function refreshCharactersList() {
  charactersContainer.innerHTML = ""
  getList('/api/character')
  .then(ch => {
    if (ch && ch.length > 0) {
      ch.forEach(n => {
        charactersContainer.appendChild( createItem(n, () => {
          deleteEntry('/api/character/' + n._id)
          .then(() => refresh())
          .catch(e => errorP = 'Error deleting Character.' + e)
        }));
      });
    }
  })
  .catch(e => errorP = 'Error fetching Characters.' + e)
}


function refreshNumbersList() {
  numbersContainer.innerHTML = ""
  getList('/api/number')
  .then(ns => {
    if (ns && ns.length > 0) {
      ns.forEach(n => {
        numbersContainer.appendChild( createItem(n, () => {
          deleteEntry('/api/number/' + n._id)
          .then(() => refresh())
          .catch(e => errorP = 'Error deleting Number.' + e)
        }));
      });
    }
  })
  .catch(e => errorP = 'Error fetching Numbers.' + e)
}

function refreshTextsList() {
  textsContainer.innerHTML = ""
  getList('/api/text')
  .then(ts => {
    if (ts && ts.length > 0) {
      ts.forEach(t => {
        textsContainer.appendChild( createItem(t, () => {
          deleteEntry('/api/text/' + t._id)
          .then(() => refresh())
          .catch(e => errorP = 'Error deleting Number.' + e)
        }));
      });
    }
  })
  .catch(e => errorP = 'Error fetching Numbers.' + e)
}


/**     UI Helpers     */

function createItem(number, cb) {
  const item = d.createElement('li');
  item.innerText = number.value

  const buttonDelete = d.createElement('button');
  buttonDelete.innerText = 'Delete';

  buttonDelete.addEventListener('click', () => {
      cb(number.id)
      .then(() => refresh())
      .catch(e => {
        errorP = 'Error deleting Category';
      });
  });

  item.append( buttonDelete );

  return item;
}

/*     API Calls    */

async function deleteEntry(api) {
  const tm = setTimeout(() => {
    errorP = 'There appears to be an error comunicating with API.';
  }, 3000);
  try {
    await fetch(api, { method: 'DELETE' });
    return;
  } catch (e) {
    throw e;
  } finally {
    clearTimeout(tm);
  }
}

async function postInput(data) {
  const tm = setTimeout(() => {
    errorP = 'There appears to be an error comunicating with API.';
  }, 3000);
  try {
    await fetch('/api/input', 
    {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ value: data })
    }
    );
    return;
  } catch (e) {
    throw e;
  } finally {
    clearTimeout(tm);
  }
}

async function getList(url) {
  const tm = setTimeout(() => {
    errorP = 'There appears to be an error comunicating with API.';
  }, 3000);
  try {
    const resp = await fetch(url, { method: 'GET' });
    const data = await resp.json();
    return data;
  } catch (e) {
    throw e;
  } finally {
    clearTimeout(tm);
  }
}
