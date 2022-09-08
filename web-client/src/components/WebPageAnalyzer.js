import {useState} from 'react';
import Typography from '@mui/material/Typography';
import { Button, Grid, TextField, Box } from '@mui/material';
import { getAnalyze } from '../api/webpage-analyzer';

const WebPageAnalyzer = () => {

    const [url, setURL] = useState("");

    const analyzeWebpage = () => {
        getAnalyze(url)
    }

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
                    HTML version: 5.1
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Page title: xxxxx
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    How many headings: xxxxx
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    How many external and internal links: xxxxx
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Are there any inaccessible links, if so how many: Yes, 5
                </Typography>
                <Typography variant="subtitle1" gutterBottom align='left'>
                    Is this page containing a login form: Yes
                </Typography>
            </Box>
        </Grid>
      </Grid>
    );
};

export default WebPageAnalyzer;//shadiya usman ashkar