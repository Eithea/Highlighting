import { Link } from "react-router-dom";

function Home() {
    return (
        <div>
            <input />
            <Link to="/archive-viewer">
                <button> 아카이브 뷰어 </button>
            </Link>
            <Link to="/clip-collections">
                <button> 클립 모음 </button>
            </Link>
        </div>
    );
}

export default Home;
