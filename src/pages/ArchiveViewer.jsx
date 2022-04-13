import "./ArchiveViewer.scss"

import Header from "../components/header/Header";

function ArchiveViewer() {
    return (
        <>
            <Header/>
            <div class = 'Cover-ArchiveViewer'>
                <div class = 'Cover-Viewers'>
                    <div class = 'Cover-Broadcast'>
                        <div class = 'Cover-Video'>
                            <div class = 'Comp-VideoPlayer'>(비디오 플레이어)</div>
                            <div class = 'Comp-VideoController'>(비디오 제어)</div>
                        </div>
                        <div class = 'Cover-Chat'>
                        <div class = 'Comp-ChatViewer'>(채팅 뷰어)</div>
                        <div class = 'Comp-ChatAnalyzer'>(채팅 분석)</div>
                        </div>
                    </div>
                    <div class = 'Cover-Data'>
                        <div class = 'Comp-DataIndex'>(데이터 인덱스)</div>
                        <div class = 'Comp-DataChart'>(데이터 차트)</div>
                    </div>
                </div>
                <div class = 'Cover-Sidebar'>
                    <div class = 'Comp-ViewChanger'>(사이드바 뷰 체인저)</div>
                    <div class = 'Comp-ClipList'>(클립 리스트)</div>
                    <div class = 'Comp-ArchiveList'>(아카이브 리스트)</div>
                </div>
            </div>
        </>
    );
}

export default ArchiveViewer;
