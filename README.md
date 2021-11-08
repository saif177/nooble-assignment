<center>
# nooble-assignment


technology stack -  <br />
  database - Postgresql <br />
  langauge - Golang <br />
  additional - 3rd parties git libraries <br />
              &emsp; -github.com/gorilla/mux <br />
	       &emsp;      -github.com/jinzhu/gorm <br />
	       &emsp;      -github.com/jinzhu/gorm/dialects/postgres <br />

the assignment is partially done and work is still in progress <br />

there are two table , one is of Users table and another one is AudioFile. <br />

there are total 4 routes  <br />
	&emsp;<b>(get) </b> - get all audio list.  <br />
	&emsp; <b>(get) </b> - get invidual audio basad on id. <br />
	&emsp; <b>(post) </b> - upload new audio.  <br />
        &emsp; <b> (delete) </b>- delete invidual audio based on id.  <br />


<h2>  planned designed process for storing audio files</h2> <br />
	<b> setp - 1 </b>   get all the details like title , description etc. <br />
	<b> step - 2 </b>   store file in AWS s3 bucket and get file path <br />
	<b> step - 3 </b>   pass file path as file along with remain details to store in <b>postgresql</b>. <br />
</center> 

