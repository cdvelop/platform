const URL_HTTP =   window.location.protocol+"//" + window.location.host;


function NewRequest(props) {
    const { method, contentType, endpoint, data, response } = props;
    const content_type = contentType || 'application/json';
    
    fetch(URL_HTTP + endpoint, {
      method,
      headers: {
        'Content-Action': content_type
      },
      body: content_type === 'application/json' ? JSON.stringify(data) : data
    })
    .then(res => {
      if (!res.ok) {
        throw new Error('Network response was not ok');
      }
      return res.json();
    })
    .then(json => response(json))
    .catch(error => {
      ShowMessageToUser({"Action":"error","Message":error.toString()})
    });
  }