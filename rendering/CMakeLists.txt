project(dywoqgame_rendering)

cmake_minimum_required(VERSION 4.1.1)

set(CMAKE_CXX_STANDARD 20)
file(GLOB SOURCES src/*.cxx)

include(FetchContent)

# raylib
FetchContent_Declare(
	raylib
	GIT_REPOSITORY https://github.com/raysan5/raylib.git
	GIT_TAG		   master
)
FetchContent_MakeAvailable(raylib)

add_library(${PROJECT_NAME} STATIC
	${SOURCES}
)

target_link_libraries(${PROJECT_NAME} PUBLIC raylib)
target_include_directories(${PROJECT_NAME} PUBLIC ${raylib_SOURCE_DIR}/src)
