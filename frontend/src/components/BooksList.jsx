// src/components/BooksList.js
import styled from 'styled-components';

export const MainContainer = styled.section`
width: 75vw;
height: 75vh;
border: 2px groove lime;
display: flex;
justify-content:center;
align-items:center;
margin: auto;  
border-radius: 1.5rem;
padding:.5rem;
`;
export const Wrapper = styled.article`
width: 100%;
height: 100%;
>h2{
  text-align:center;
}
`;
export const Wrap = styled.div`
display:flex;
justify-content: space-evenly;
align-items:center;
padding:.5rem;
`;
export const BooksListContainer = styled.div`
  padding:1rem;
  border: 1px solid lime;
  background:salmon;
  `;
  
  export const BookListItem = styled.li`
  padding: 10px;
  margin: 5px 0;
  border: 1px solid lime;
  border-radius: 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-right: 35px;

  button {
    margin-left: 10px;
  }
`;


