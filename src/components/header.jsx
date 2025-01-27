import './styles/header.css';

export function Header() {
    return(
        <header className='header'>
            <div className='container'>
                <h1 className='logo'><img src='https://cdn-icons-png.flaticon.com/512/116/116392.png' className='gift'></img>Gift</h1>
                <button className='fav'>Favorites<img src='https://cdn-icons-png.flaticon.com/512/158/158722.png' className='heart'></img></button>
            </div>
        </header>
    )
}

