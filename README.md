# Creatly Leads API
## Service that collects info from landing page's lead form & saves it to CRM (Trello Board)

### Prerequisites

- go 1.16

### Build & Run

1. Set env variables
```
export TRELLO_LIST_NAME="Leads"
export TRELLO_BOARD_ID=TqW02k3K
export TRELLO_API_TOKEN=<your token>
export TRELLO_API_KEY=<your key>
```
2. Run app
```
make run
```

### API
<b>POST /leads</b>
```json
{
    "firstname": "Vasya",
    "lastname": "Pupkin",
    "email": "vasyapupkin@gmail.com",
    "phone": "+380996969228",
    "source": "pdf",
    "niche": "IT"
}
```