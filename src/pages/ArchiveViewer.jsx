import { Link } from "react-router-dom";

function ArchiveViewer() {
    return (
        <div>
            아카이브 뷰어
            <Link to="/clip-collections">
                <button> 클립 모음 </button>
            </Link>
            <Link to="/">
                <button> 홈 </button>
            </Link>
        </div>
    );
}

export default ArchiveViewer;
