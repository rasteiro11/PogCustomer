import {useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {Card, Client} from "../Client"
// import {Card as CardComponent} from "./Card"

var cardIdx = 0
const initialCards: Card[] = []
export default function Play(props: {client: Client}) {
  let client = props.client as Client
  let navigate = useNavigate()
  const [cards, setCards] = useState(initialCards)
  const [card, setCard] = useState(new Card(0, "", ""))

  useEffect(() => {
    if (client.token.Token == "") {
      navigate("/")
      return
    }
    // client.listCards().then((list) => {
    //   setCards(list.Data)
    // 
    //   if (list.Data !== null && list.Data.length !== 0) {
    //     setCard(list.Data[cardIdx])
    //   }
    // })
  }, [])

  function showCard() {
    if (cards != null) {
      return (
        <div className="d-flex justify-content-center" >
          <div className="container">
              <div className="container">
                <div className="row">
                  <a onClick={() => {
                    if (cardIdx < cards.length - 1) {
                      cardIdx++
                      setCard(cards[cardIdx])
                    }
                  }} className="btn btn-success col">Next</a>
                  <a onClick={() => {
                    if (cardIdx > 0) {
                      cardIdx--
                      setCard(cards[cardIdx])
                    }
                  }} className="btn btn-danger col">Previous</a>
                </div>
              </div>
            </div>
          </div>
      )
    }
  }

  return (
    <>
      {showCard()}
    </>
  )
}


