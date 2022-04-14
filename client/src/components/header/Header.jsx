import { Link } from "react-router-dom";

function Header() {
    return (
        <div>
            헤더
            <Link to="/">
                <button> 홈 </button>
            </Link>
            <Link to="/archive-viewer">
                <button> 아카이브 뷰어 </button>
            </Link>
            <Link to="/clip-collections">
                <button> 클립 모음 </button>
            </Link>
        </div>
    );
}

export default Header;