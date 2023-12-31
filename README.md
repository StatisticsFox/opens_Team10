<h2 align="center"> Open Software Team 10 </h2>
<h2 align="center">  최지혁 - 유윤정 - 보꾸옥안 </h2>
  <div align="center">
    <span><img src="https://img.shields.io/badge/Kubernetes-282C34?logo=kubernetes&logoColor=0000FF" alt="Kubernetes logo" title="Kubernetes" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/Docker-282C34?logo=docker&logoColor=2496ED" alt="Docker logo" title="Docker" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/Jenkins-282C34?logo=jenkins&logoColor=EEEEEE" alt="Jenkins logo" title="Docker" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/Github-282C34?logo=github&logoColor=181717" alt="github logo" title="github" height="25" /></span>&nbsp;</div>
  <div align="center">
    <span><img src="https://img.shields.io/badge/Visual studio code-282C34?logo=visualstudiocode&logoColor=2496ED" alt="visualstudiocode logo" title="visualstudiocode" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/Git-282C34?logo=git&logoColor=F05032" alt="git logo" title="git" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/MacOs-282C34?logo=macos&logoColor=000000" alt="macos logo" title="macos" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/Windows 11-282C34?logo=windows11&logoColor=0078D6" alt="windows11 logo" title="windows11" height="25" /></span>&nbsp;
    <span><img src="https://img.shields.io/badge/MySQL-282C34?logo=mysql&logoColor=4479A1" alt="mysql logo" title="mysql" height="25" /></span>&nbsp;
  </div>
    
## What is Workflows?
 Kubernetes (Docker Desktop 사용 / 실습이 가능하면 Minikube 사용 가능)로 클러스터 구성
- 마스터 노드(1) + 워커 노드(4개 이상, 팀원 수 만큼) + 2개 이상의 서비스 및 하위 POD 구성 
- 서비스는 젠킨스, GITHUB, 웹 서비스(공개 되어 있는 서비스 활용 가능)
- 웹 서비스의 입력 내용이 GITHUB / DB 등에 저장되어야 함
- 웹 서비스는 팀별로 원하거나 구현했던 서비스 적용
- 물리적 워커노드 2개 이상이며, 하나의 Windows PC에 WSL과 버추얼 박스로 구성 가능
  
<h2 align="left"> 프로젝트 구성도  </h2>
    <span><img src="https://github.com/mitsumi73/kubectl/blob/main/KakaoTalk_Photo_2023-12-22-17-48-03.png"/></span>&nbsp;
<h2 align="left"> Minikube & Kubectl </h2>
    <span><img src="https://github.com/mitsumi73/kubectl/blob/main/minikube.png" height="300"/></span>&nbsp;
    <span><img src="https://github.com/mitsumi73/kubectl/blob/main/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-12-21%20%EC%98%A4%EC%A0%84%2010.50.33.png" height="300"/></span>&nbsp;
<h2 align="left"> Web서비스 Deployment </h2>
    <span><img src="https://github.com/mitsumi73/kubectl/blob/main/2.png" height="300"/></span>&nbsp;
    <span><img src="https://github.com/mitsumi73/kubectl/blob/main/KakaoTalk_Image_2023-12-21-22-43-09.jpeg" height="320"/></span>&nbsp;
