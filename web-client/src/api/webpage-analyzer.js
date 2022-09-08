import Axios from 'axios'

export async function getAnalyze( 
  url
) {

  const res = await Axios.get(`/api/v1/analyze-webpage`, {
    params:{ 
      url
     }
  });
  console.log(res);
  return res
}