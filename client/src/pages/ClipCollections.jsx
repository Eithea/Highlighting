import { Link } from "react-router-dom";

function ClipCollections() {
    return (
        <div>
            클립 모음<br/>
            <Link to="/">
                <button> 홈 </button>
            </Link>
            <Link to="/archive-viewer">
                <button> 아카이브 뷰어 </button>
            </Link>
        </div>
    );
}

export default ClipCollections;
