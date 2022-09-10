import {useEffect, useState} from 'react';
import Typography from '@mui/material/Typography';
import { Button, Grid, TextField, Box } from '@mui/material';
import { getAnalyze } from '../api/webpage-analyzer';

const WebPageAnalyzer = () => {

    const [url, setURL] = useState("");
    const [scrapedData, setScrapedData] = useState({});

    const analyzeWebpage = async () => {
        let res = await getAnalyze(url)
        setScrapedData(res.data.data);
    }

    useEffect(()=>{
        console.log(scrapedData);
    },[scrapedData])

    return (
        <Grid container spacing={2} 
            sx={{
                marginTop: 2,
                marginLeft: 2
            }}
        >
        <Grid container item xs={8} direction="column">
            <TextField id="url" label="URL to be analyzed" variant="outlined" 
            value={url} 
            onClick={() => setURL("")}
            onChange={(e) => setURL(e.target.value)}
            />
        </Grid>
        <Grid container item xs={4} >
            <Button variant="contained"
            onClick={() => analyzeWebpage()}
            >Analyze</Button>
        </Grid>
        <Grid contriner item>
            <Box
            sx={{
                marginTop: 2,
                marginLeft: 2,
                width: '100%', 
                maxWidth: 500
            }}
            >
                <Typography variant="h6" gutterBottom align='left'>
                    Analyzation report
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                {scrapedData.Version}
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Page title: {scrapedData.Title}
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Headings count level wise: 
                    {` H1:${scrapedData.H1} , H2:${scrapedData.H2} , H3:${scrapedData.H3} , H4:${scrapedData.H4} , 
                    H5:${scrapedData.H5} , H6:${scrapedData.H6}`}
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    External and internal links: {` External:${scrapedData.ExternalLink} , Internal:${scrapedData.InternalLink} `}
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Inaccessible links: 
                    {
                        scrapedData.ExternalDeadLink+scrapedData.InternalDeadIdLink+scrapedData.InternalDeadPathLink > 0 ? "Yes," : "No"
                    }
                    {
                        scrapedData.ExternalDeadLink+scrapedData.InternalDeadIdLink+scrapedData.InternalDeadPathLink > 0 ?
                        ` External: ${scrapedData.ExternalDeadLink} , Internal ID: ${scrapedData.InternalDeadIdLink} , Internal Path: ${scrapedData.InternalDeadPathLink} ` : ""
                    }
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Is this page containing a login form: { scrapedData.IsWithLogin ? "Yes" : "No" }
                </Typography>
            </Box>
        </Grid>
      </Grid>
    );
};

export default WebPageAnalyzer;//shadiya usman ashkar